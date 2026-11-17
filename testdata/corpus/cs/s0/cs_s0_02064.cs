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

    public delegate void Notifier19250(int value);

    public sealed class Source19250
    {
        public event Notifier19250 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier19251(int value);

    public sealed class Source19251
    {
        public event Notifier19251 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard19252(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 18;
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
    public static int Guard19253(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 25;
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
