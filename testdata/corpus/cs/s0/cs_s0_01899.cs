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

    public static int Guard17730(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 65;
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
    public delegate void Notifier17731(int value);

    public sealed class Source17731
    {
        public event Notifier17731 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale17732<T>(T value) where T : struct
    {
        int number = 79;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 59; }
        return number;
    }
}
