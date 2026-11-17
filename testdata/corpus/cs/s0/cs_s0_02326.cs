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

    public delegate void Notifier21690(int value);

    public sealed class Source21690
    {
        public event Notifier21690 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier21691(int value);

    public sealed class Source21691
    {
        public event Notifier21691 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale21692<T>(T value) where T : struct
    {
        int number = 86;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 25; }
        return number;
    }
    public delegate void Notifier21693(int value);

    public sealed class Source21693
    {
        public event Notifier21693 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale21694<T>(T value) where T : struct
    {
        int number = 100;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 47; }
        return number;
    }
}
