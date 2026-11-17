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

    public delegate void Notifier18810(int value);

    public sealed class Source18810
    {
        public event Notifier18810 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale18811<T>(T value) where T : struct
    {
        int number = 21;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 11; }
        return number;
    }
    public static int Scale18812<T>(T value) where T : struct
    {
        int number = 28;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 22; }
        return number;
    }
    public delegate void Notifier18813(int value);

    public sealed class Source18813
    {
        public event Notifier18813 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale18814<T>(T value) where T : struct
    {
        int number = 42;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 44; }
        return number;
    }
    public static int Guard18815(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 49;
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
