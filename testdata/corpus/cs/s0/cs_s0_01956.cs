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

    public delegate void Notifier18240(int value);

    public sealed class Source18240
    {
        public event Notifier18240 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard18241(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 56;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
    public delegate void Notifier18242(int value);

    public sealed class Source18242
    {
        public event Notifier18242 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale18243<T>(T value) where T : struct
    {
        int number = 70;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 61; }
        return number;
    }
}
