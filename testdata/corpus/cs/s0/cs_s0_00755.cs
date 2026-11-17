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

    public interface IHandler7080
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope7080(string Label, int Payload);

    public sealed class Handler7080 : IHandler7080
    {
        public void Handle(int value) { }
    }
    public interface IHandler7081
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope7081(string Label, int Payload);

    public sealed class Handler7081 : IHandler7081
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier7082(int value);

    public sealed class Source7082
    {
        public event Notifier7082 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard7083(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 15;
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
