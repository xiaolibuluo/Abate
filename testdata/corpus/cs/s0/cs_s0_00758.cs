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

    public interface IHandler7100
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope7100(string Label, int Payload);

    public sealed class Handler7100 : IHandler7100
    {
        public void Handle(int value) { }
    }
    public interface IHandler7101
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope7101(string Label, int Payload);

    public sealed class Handler7101 : IHandler7101
    {
        public void Handle(int value) { }
    }
    public static int Guard7102(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 34;
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
    public interface IHandler7103
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope7103(string Label, int Payload);

    public sealed class Handler7103 : IHandler7103
    {
        public void Handle(int value) { }
    }
}
