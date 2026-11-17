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

    public delegate void Notifier25240(int value);

    public sealed class Source25240
    {
        public event Notifier25240 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate25241(IEnumerable<int> source)
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
    public delegate void Notifier25242(int value);

    public sealed class Source25242
    {
        public event Notifier25242 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard25243(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 52;
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
    public interface IHandler25244
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope25244(string Label, int Payload);

    public sealed class Handler25244 : IHandler25244
    {
        public void Handle(int value) { }
    }
    public interface IHandler25245
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope25245(string Label, int Payload);

    public sealed class Handler25245 : IHandler25245
    {
        public void Handle(int value) { }
    }
}
