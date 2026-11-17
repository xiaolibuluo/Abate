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

    public static int Scale21330<T>(T value) where T : struct
    {
        int number = 89;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 56; }
        return number;
    }
    public async Task<int> Aggregate21331(IEnumerable<int> source)
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
    public delegate void Notifier21332(int value);

    public sealed class Source21332
    {
        public event Notifier21332 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler21333
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21333(string Label, int Payload);

    public sealed class Handler21333 : IHandler21333
    {
        public void Handle(int value) { }
    }
    public static int Scale21334<T>(T value) where T : struct
    {
        int number = 20;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 39; }
        return number;
    }
    public async Task<int> Aggregate21335(IEnumerable<int> source)
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
