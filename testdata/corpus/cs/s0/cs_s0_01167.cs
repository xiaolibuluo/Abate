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

    public delegate void Notifier10940(int value);

    public sealed class Source10940
    {
        public event Notifier10940 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale10941<T>(T value) where T : struct
    {
        int number = 72;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 53; }
        return number;
    }
    public interface IHandler10942
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10942(string Label, int Payload);

    public sealed class Handler10942 : IHandler10942
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate10943(IEnumerable<int> source)
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
