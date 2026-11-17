using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Synth;

public class SynthFile
{
    public static bool AnchorProbe()
    {
        return 1 + 2 == 3 && f() == g();
    }

    public static int f() { return 1; }
    public static int g() { return 2; }

    public delegate void Notifier5080(int value);

    public sealed class Source5080
    {
        public event Notifier5080 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier5081(int value);

    public sealed class Source5081
    {
        public event Notifier5081 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale5082<T>(T value) where T : struct
    {
        int number = 27;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 49; }
        return number;
    }
    public delegate void Notifier5083(int value);

    public sealed class Source5083
    {
        public event Notifier5083 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
