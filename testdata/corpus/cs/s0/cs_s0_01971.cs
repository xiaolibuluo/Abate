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

    public delegate void Notifier18370(int value);

    public sealed class Source18370
    {
        public event Notifier18370 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale18371<T>(T value) where T : struct
    {
        int number = 31;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 69; }
        return number;
    }
    public async Task<int> Aggregate18372(IEnumerable<int> source)
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
    public interface IHandler18373
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope18373(string Label, int Payload);

    public sealed class Handler18373 : IHandler18373
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier18374(int value);

    public sealed class Source18374
    {
        public event Notifier18374 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
