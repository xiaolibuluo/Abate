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

    public static int Guard10870(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 71;
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
    public delegate void Notifier10871(int value);

    public sealed class Source10871
    {
        public event Notifier10871 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler10872
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10872(string Label, int Payload);

    public sealed class Handler10872 : IHandler10872
    {
        public void Handle(int value) { }
    }
}
