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

    public delegate void Notifier8940(int value);

    public sealed class Source8940
    {
        public event Notifier8940 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard8941(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 91;
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
    public interface IHandler8942
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8942(string Label, int Payload);

    public sealed class Handler8942 : IHandler8942
    {
        public void Handle(int value) { }
    }
}
