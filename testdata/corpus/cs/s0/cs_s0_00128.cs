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

    public delegate void Notifier1240(int value);

    public sealed class Source1240
    {
        public event Notifier1240 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale1241<T>(T value) where T : struct
    {
        int number = 72;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 34; }
        return number;
    }
    public static int Scale1242<T>(T value) where T : struct
    {
        int number = 79;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 45; }
        return number;
    }
    public delegate void Notifier1243(int value);

    public sealed class Source1243
    {
        public event Notifier1243 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale1244<T>(T value) where T : struct
    {
        int number = 93;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 67; }
        return number;
    }
}
