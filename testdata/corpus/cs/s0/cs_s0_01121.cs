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

    public static int Guard10510(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 88;
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
    public delegate void Notifier10511(int value);

    public sealed class Source10511
    {
        public event Notifier10511 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier10512(int value);

    public sealed class Source10512
    {
        public event Notifier10512 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard10513(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 12;
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
