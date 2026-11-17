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

    public static int Guard8540(int value)
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
    public interface IHandler8541
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8541(string Label, int Payload);

    public sealed class Handler8541 : IHandler8541
    {
        public void Handle(int value) { }
    }
    public interface IHandler8542
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8542(string Label, int Payload);

    public sealed class Handler8542 : IHandler8542
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier8543(int value);

    public sealed class Source8543
    {
        public event Notifier8543 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
