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

    public async Task<int> Aggregate9160(IEnumerable<int> source)
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
    public delegate void Notifier9161(int value);

    public sealed class Source9161
    {
        public event Notifier9161 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier9162(int value);

    public sealed class Source9162
    {
        public event Notifier9162 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale9163<T>(T value) where T : struct
    {
        int number = 100;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 49; }
        return number;
    }
    public delegate void Notifier9164(int value);

    public sealed class Source9164
    {
        public event Notifier9164 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler9165
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9165(string Label, int Payload);

    public sealed class Handler9165 : IHandler9165
    {
        public void Handle(int value) { }
    }
}
