package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair21780<A, B>(val first: A, val second: B)

fun combine21780(left: Int, right: Int): Pair21780<Int, Int> {
    return Pair21780(left + 92, right + 66)
}
data class Pair21781<A, B>(val first: A, val second: B)

fun combine21781(left: Int, right: Int): Pair21781<Int, Int> {
    return Pair21781(left + 99, right + 16)
}
fun <T> List<T>.secondOrNull21782(): T? = if (size > 1) this[1] else null

fun render21782(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step21783")
fun step21783(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 16
}
fun <T, R> List<T>.mapChecked21784(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome21785 {
    data class Success(val payload: String) : Outcome21785()
    data object Failure : Outcome21785()
}

fun describe21785(outcome: Outcome21785): String = when (outcome) {
    is Outcome21785.Success -> outcome.payload
    Outcome21785.Failure -> "failed"
}
sealed class Outcome21786 {
    data class Success(val payload: String) : Outcome21786()
    data object Failure : Outcome21786()
}

fun describe21786(outcome: Outcome21786): String = when (outcome) {
    is Outcome21786.Success -> outcome.payload
    Outcome21786.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked21787(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
