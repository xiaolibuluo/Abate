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

    public delegate void Notifier19880(int value);

    public sealed class Source19880
    {
        public event Notifier19880 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate19881(IEnumerable<int> source)
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
    public delegate void Notifier19882(int value);

    public sealed class Source19882
    {
        public event Notifier19882 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale19883<T>(T value) where T : struct
    {
        int number = 68;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 50; }
        return number;
    }
    public interface IHandler19884
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19884(string Label, int Payload);

    public sealed class Handler19884 : IHandler19884
    {
        public void Handle(int value) { }
    }
    public static int Scale19885<T>(T value) where T : struct
    {
        int number = 82;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 11; }
        return number;
    }
}
