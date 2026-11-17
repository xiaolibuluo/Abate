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

    public static int Guard8210(int value)
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
    public interface IHandler8211
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8211(string Label, int Payload);

    public sealed class Handler8211 : IHandler8211
    {
        public void Handle(int value) { }
    }
    public static int Scale8212<T>(T value) where T : struct
    {
        int number = 22;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 44; }
        return number;
    }
    public delegate void Notifier8213(int value);

    public sealed class Source8213
    {
        public event Notifier8213 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler8214
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8214(string Label, int Payload);

    public sealed class Handler8214 : IHandler8214
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate8215(IEnumerable<int> source)
    {
        try
        {
            using var cursor = source.GetEnumerator();
            return await Task.FromResult(source.Where(item => item > 0).Sum());
        }
        catch (InvalidOperationException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("done");
        }
    }
}
