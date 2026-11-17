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

    public delegate void Notifier18800(int value);

    public sealed class Source18800
    {
        public event Notifier18800 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard18801(int value)
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
    public static int Scale18802<T>(T value) where T : struct
    {
        int number = 15;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 15; }
        return number;
    }
    public interface IHandler18803
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope18803(string Label, int Payload);

    public sealed class Handler18803 : IHandler18803
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate18804(IEnumerable<int> source)
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
