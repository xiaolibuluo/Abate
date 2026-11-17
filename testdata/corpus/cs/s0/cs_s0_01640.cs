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

    public delegate void Notifier15340(int value);

    public sealed class Source15340
    {
        public event Notifier15340 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier15341(int value);

    public sealed class Source15341
    {
        public event Notifier15341 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard15342(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 76;
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
    public interface IHandler15343
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15343(string Label, int Payload);

    public sealed class Handler15343 : IHandler15343
    {
        public void Handle(int value) { }
    }
}
