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

    public static int Guard12480(int value)
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
    public delegate void Notifier12481(int value);

    public sealed class Source12481
    {
        public event Notifier12481 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard12482(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 44;
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
    public interface IHandler12483
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12483(string Label, int Payload);

    public sealed class Handler12483 : IHandler12483
    {
        public void Handle(int value) { }
    }
}
