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

    public interface IHandler5140
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5140(string Label, int Payload);

    public sealed class Handler5140 : IHandler5140
    {
        public void Handle(int value) { }
    }
    public interface IHandler5141
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5141(string Label, int Payload);

    public sealed class Handler5141 : IHandler5141
    {
        public void Handle(int value) { }
    }
    public static int Guard5142(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 8;
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
    public delegate void Notifier5143(int value);

    public sealed class Source5143
    {
        public event Notifier5143 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
