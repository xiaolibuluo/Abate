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

    public static int Guard4180(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 7;
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
    public delegate void Notifier4181(int value);

    public sealed class Source4181
    {
        public event Notifier4181 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale4182<T>(T value) where T : struct
    {
        int number = 21;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 29; }
        return number;
    }
    public static int Guard4183(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 28;
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
}
