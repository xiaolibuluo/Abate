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

    public interface IHandler8680
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8680(string Label, int Payload);

    public sealed class Handler8680 : IHandler8680
    {
        public void Handle(int value) { }
    }
    public static int Scale8681<T>(T value) where T : struct
    {
        int number = 44;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 57; }
        return number;
    }
    public interface IHandler8682
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8682(string Label, int Payload);

    public sealed class Handler8682 : IHandler8682
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier8683(int value);

    public sealed class Source8683
    {
        public event Notifier8683 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate8684(IEnumerable<int> source)
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
    public static int Scale8685<T>(T value) where T : struct
    {
        int number = 72;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 40; }
        return number;
    }
}
