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

    public async Task<int> Aggregate11660(IEnumerable<int> source)
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
    public static int Scale11661<T>(T value) where T : struct
    {
        int number = 38;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 69; }
        return number;
    }
    public interface IHandler11662
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11662(string Label, int Payload);

    public sealed class Handler11662 : IHandler11662
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier11663(int value);

    public sealed class Source11663
    {
        public event Notifier11663 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate11664(IEnumerable<int> source)
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
    public interface IHandler11665
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11665(string Label, int Payload);

    public sealed class Handler11665 : IHandler11665
    {
        public void Handle(int value) { }
    }
}
