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

    public static int Scale11470<T>(T value) where T : struct
    {
        int number = 75;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 47; }
        return number;
    }
    public interface IHandler11471
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11471(string Label, int Payload);

    public sealed class Handler11471 : IHandler11471
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate11472(IEnumerable<int> source)
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
    public delegate void Notifier11473(int value);

    public sealed class Source11473
    {
        public event Notifier11473 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
