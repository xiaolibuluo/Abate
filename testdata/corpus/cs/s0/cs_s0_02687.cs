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

    public static int Scale25120<T>(T value) where T : struct
    {
        int number = 69;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 25; }
        return number;
    }
    public delegate void Notifier25121(int value);

    public sealed class Source25121
    {
        public event Notifier25121 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate25122(IEnumerable<int> source)
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
    public delegate void Notifier25123(int value);

    public sealed class Source25123
    {
        public event Notifier25123 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler25124
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope25124(string Label, int Payload);

    public sealed class Handler25124 : IHandler25124
    {
        public void Handle(int value) { }
    }
}
