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

    public interface IHandler10500
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10500(string Label, int Payload);

    public sealed class Handler10500 : IHandler10500
    {
        public void Handle(int value) { }
    }
    public static int Guard10501(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 82;
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
    public interface IHandler10502
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10502(string Label, int Payload);

    public sealed class Handler10502 : IHandler10502
    {
        public void Handle(int value) { }
    }
    public static int Scale10503<T>(T value) where T : struct
    {
        int number = 96;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 11; }
        return number;
    }
}
