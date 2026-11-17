package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome14490 {
    data class Success(val payload: String) : Outcome14490()
    data object Failure : Outcome14490()
}

fun describe14490(outcome: Outcome14490): String = when (outcome) {
    is Outcome14490.Success -> outcome.payload
    Outcome14490.Failure -> "failed"
}
object Registry14491 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14491 = "registry-31"
    }
}
data class Pair14492<A, B>(val first: A, val second: B)

fun combine14492(left: Int, right: Int): Pair14492<Int, Int> {
    return Pair14492(left + 38, right + 48)
}
fun <T, R> List<T>.mapChecked14493(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step14494")
fun step14494(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 52
}
fun <T, R> List<T>.mapChecked14495(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair14496<A, B>(val first: A, val second: B)

fun combine14496(left: Int, right: Int): Pair14496<Int, Int> {
    return Pair14496(left + 66, right + 31)
}
@JvmName("step14497")
fun step14497(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 73
}
fun <T, R> List<T>.mapChecked14498(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
