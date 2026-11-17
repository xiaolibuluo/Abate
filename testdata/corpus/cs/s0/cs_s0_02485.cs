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

    public static int Guard23200(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 95;
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
    public interface IHandler23201
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23201(string Label, int Payload);

    public sealed class Handler23201 : IHandler23201
    {
        public void Handle(int value) { }
    }
    public interface IHandler23202
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23202(string Label, int Payload);

    public sealed class Handler23202 : IHandler23202
    {
        public void Handle(int value) { }
    }
    public interface IHandler23203
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23203(string Label, int Payload);

    public sealed class Handler23203 : IHandler23203
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier23204(int value);

    public sealed class Source23204
    {
        public event Notifier23204 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
