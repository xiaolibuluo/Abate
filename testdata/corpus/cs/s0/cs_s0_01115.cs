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

    public interface IHandler10450
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10450(string Label, int Payload);

    public sealed class Handler10450 : IHandler10450
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier10451(int value);

    public sealed class Source10451
    {
        public event Notifier10451 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard10452(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 24;
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
