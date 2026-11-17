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

    public static int Scale12620<T>(T value) where T : struct
    {
        int number = 18;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 59; }
        return number;
    }
    public static int Scale12621<T>(T value) where T : struct
    {
        int number = 25;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 9; }
        return number;
    }
    public async Task<int> Aggregate12622(IEnumerable<int> source)
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
    public delegate void Notifier12623(int value);

    public sealed class Source12623
    {
        public event Notifier12623 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale12624<T>(T value) where T : struct
    {
        int number = 46;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 42; }
        return number;
    }
    public interface IHandler12625
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12625(string Label, int Payload);

    public sealed class Handler12625 : IHandler12625
    {
        public void Handle(int value) { }
    }
    public struct Point12626
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 60;
    }

    public static class Grid12626
    {
        public static Point12626 Make(int left, int right)
        {
            return new Point12626 { X = left, Y = right };
        }
    }
}
