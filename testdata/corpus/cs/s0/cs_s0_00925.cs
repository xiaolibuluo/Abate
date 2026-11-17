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

    public interface IHandler8660
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8660(string Label, int Payload);

    public sealed class Handler8660 : IHandler8660
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier8661(int value);

    public sealed class Source8661
    {
        public event Notifier8661 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate8662(IEnumerable<int> source)
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
    public interface IHandler8663
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8663(string Label, int Payload);

    public sealed class Handler8663 : IHandler8663
    {
        public void Handle(int value) { }
    }
    public static int Scale8664<T>(T value) where T : struct
    {
        int number = 39;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 15; }
        return number;
    }
    public delegate void Notifier8665(int value);

    public sealed class Source8665
    {
        public event Notifier8665 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
