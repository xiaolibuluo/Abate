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

    public delegate void Notifier6300(int value);

    public sealed class Source6300
    {
        public event Notifier6300 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate6301(IEnumerable<int> source)
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
    public delegate void Notifier6302(int value);

    public sealed class Source6302
    {
        public event Notifier6302 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale6303<T>(T value) where T : struct
    {
        int number = 68;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 60; }
        return number;
    }
    public delegate void Notifier6304(int value);

    public sealed class Source6304
    {
        public event Notifier6304 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler6305
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6305(string Label, int Payload);

    public sealed class Handler6305 : IHandler6305
    {
        public void Handle(int value) { }
    }
    public interface IHandler6306
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6306(string Label, int Payload);

    public sealed class Handler6306 : IHandler6306
    {
        public void Handle(int value) { }
    }
}
