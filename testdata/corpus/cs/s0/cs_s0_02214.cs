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

    public static int Guard20680(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 20;
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
    public static int Scale20681<T>(T value) where T : struct
    {
        int number = 27;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 39; }
        return number;
    }
    public delegate void Notifier20682(int value);

    public sealed class Source20682
    {
        public event Notifier20682 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier20683(int value);

    public sealed class Source20683
    {
        public event Notifier20683 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
