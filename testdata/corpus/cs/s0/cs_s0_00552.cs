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

    public static int Guard5170(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 33;
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
    public delegate void Notifier5171(int value);

    public sealed class Source5171
    {
        public event Notifier5171 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale5172<T>(T value) where T : struct
    {
        int number = 47;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 51; }
        return number;
    }
    public delegate void Notifier5173(int value);

    public sealed class Source5173
    {
        public event Notifier5173 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale5174<T>(T value) where T : struct
    {
        int number = 61;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 12; }
        return number;
    }
}
