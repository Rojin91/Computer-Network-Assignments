package torrent

import (
    "bufio"
    "bytes"
    "fmt"
    "reflect"
    "strconv"
)

// Unmarshal parses the bencoded data and stores the result in the value pointed to by v.
func Unmarshal(data []byte, v interface{}) error {
    r := bufio.NewReader(bytes.NewReader(data))
    return unmarshal(r, reflect.ValueOf(v))
}

func unmarshal(r *bufio.Reader, v reflect.Value) error {
    b, err := r.ReadByte()
    if err != nil {
        return err
    }

    switch {
    case b == 'i':
        return unmarshalInt(r, v)
    case b == 'l':
        return unmarshalList(r, v)
    case b == 'd':
        return unmarshalDict(r, v)
    case b >= '0' && b <= '9':
        return unmarshalString(r, v, b)
    default:
        return fmt.Errorf("invalid bencode prefix: %c", b)
    }
}

func unmarshalInt(r *bufio.Reader, v reflect.Value) error {
    var buf bytes.Buffer
    for {
        b, err := r.ReadByte()
        if err != nil {
            return err
        }
        if b == 'e' {
            break
        }
        buf.WriteByte(b)
    }
    n, err := strconv.ParseInt(buf.String(), 10, 64)
    if err != nil {
        return err
    }
    v.Elem().SetInt(n)
    return nil
}

func unmarshalString(r *bufio.Reader, v reflect.Value, prefix byte) error {
    var buf bytes.Buffer
    buf.WriteByte(prefix)
    for {
        b, err := r.ReadByte()
        if err != nil {
            return err
        }
        if b == ':' {
            break
        }
        buf.WriteByte(b)
    }
    length, err := strconv.Atoi(buf.String())
    if err != nil {
        return err
    }
    str := make([]byte, length)
    _, err = r.Read(str)
    if err != nil {
        return err
    }
    v.Elem().SetString(string(str))
    return nil
}

func unmarshalList(r *bufio.Reader, v reflect.Value) error {
    if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Slice {
        return fmt.Errorf("invalid type for list: %v", v.Kind())
    }
    v = v.Elem()
    for {
        b, err := r.ReadByte()
        if err != nil {
            return err
        }
        if b == 'e' {
            break
        }
        r.UnreadByte()
        elem := reflect.New(v.Type().Elem())
        if err := unmarshal(r, elem); err != nil {
            return err
        }
        v.Set(reflect.Append(v, elem.Elem()))
    }
    return nil
}

func unmarshalDict(r *bufio.Reader, v reflect.Value) error {
    if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Map {
        return fmt.Errorf("invalid type for dict: %v", v.Kind())
    }
    v = v.Elem()
    if v.IsNil() {
        v.Set(reflect.MakeMap(v.Type()))
    }
    for {
        b, err := r.ReadByte()
        if err != nil {
            return err
        }
        if b == 'e' {
            break
        }
        r.UnreadByte()
        key := reflect.New(reflect.TypeOf(""))
        if err := unmarshal(r, key); err != nil {
            return err
        }
        elem := reflect.New(v.Type().Elem())
        if err := unmarshal(r, elem); err != nil {
            return err
        }
        v.SetMapIndex(key.Elem(), elem.Elem())
    }
    return nil
}

// Marshal returns the bencoded encoding of v.
func Marshal(v interface{}) ([]byte, error) {
    var buf bytes.Buffer
    if err := marshal(&buf, reflect.ValueOf(v)); err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

func marshal(w *bytes.Buffer, v reflect.Value) error {
    switch v.Kind() {
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        _, err := fmt.Fprintf(w, "i%de", v.Int())
        return err
    case reflect.String:
        _, err := fmt.Fprintf(w, "%d:%s", len(v.String()), v.String())
        return err
    case reflect.Slice:
        _, err := w.WriteString("l")
        if err != nil {
            return err
        }
        for i := 0; i < v.Len(); i++ {
            if err := marshal(w, v.Index(i)); err != nil {
                return err
            }
        }
        _, err = w.WriteString("e")
        return err
    case reflect.Map:
        _, err := w.WriteString("d")
        if err != nil {
            return err
        }
        for _, key := range v.MapKeys() {
            if err := marshal(w, key); err != nil {
                return err
            }
            if err := marshal(w, v.MapIndex(key)); err != nil {
                return err
            }
        }
        _, err = w.WriteString("e")
        return err
    default:
        return fmt.Errorf("unsupported type: %v", v.Kind())
    }
}
