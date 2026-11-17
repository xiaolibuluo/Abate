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

    public async Task<int> Aggregate9510(IEnumerable<int> source)
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
    public delegate void Notifier9511(int value);

    public sealed class Source9511
    {
        public event Notifier9511 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale9512<T>(T value) where T : struct
    {
        int number = 63;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 39; }
        return number;
    }
    public interface IHandler9513
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9513(string Label, int Payload);

    public sealed class Handler9513 : IHandler9513
    {
        public void Handle(int value) { }
    }
}
