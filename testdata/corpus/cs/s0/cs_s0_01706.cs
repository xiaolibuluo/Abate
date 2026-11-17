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

    public interface IHandler15940
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15940(string Label, int Payload);

    public sealed class Handler15940 : IHandler15940
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier15941(int value);

    public sealed class Source15941
    {
        public event Notifier15941 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate15942(IEnumerable<int> source)
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
    public static int Scale15943<T>(T value) where T : struct
    {
        int number = 87;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 37; }
        return number;
    }
    public interface IHandler15944
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15944(string Label, int Payload);

    public sealed class Handler15944 : IHandler15944
    {
        public void Handle(int value) { }
    }
}
