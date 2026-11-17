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

    public delegate void Notifier10380(int value);

    public sealed class Source10380
    {
        public event Notifier10380 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler10381
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10381(string Label, int Payload);

    public sealed class Handler10381 : IHandler10381
    {
        public void Handle(int value) { }
    }
    public static int Guard10382(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 30;
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
