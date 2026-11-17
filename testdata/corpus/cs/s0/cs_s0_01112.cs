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

    public static int Guard10420(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 68;
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
    public delegate void Notifier10421(int value);

    public sealed class Source10421
    {
        public event Notifier10421 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier10422(int value);

    public sealed class Source10422
    {
        public event Notifier10422 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
