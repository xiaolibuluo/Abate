package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair13980<A, B>(val first: A, val second: B)

fun combine13980(left: Int, right: Int): Pair13980<Int, Int> {
    return Pair13980(left + 40, right + 35)
}
sealed class Outcome13981 {
    data class Success(val payload: String) : Outcome13981()
    data object Failure : Outcome13981()
}

fun describe13981(outcome: Outcome13981): String = when (outcome) {
    is Outcome13981.Success -> outcome.payload
    Outcome13981.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked13982(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step13983")
fun step13983(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 61
}
data class Pair13984<A, B>(val first: A, val second: B)

fun combine13984(left: Int, right: Int): Pair13984<Int, Int> {
    return Pair13984(left + 68, right + 18)
}
@JvmName("step13985")
fun step13985(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 75
}
fun <T> List<T>.secondOrNull13986(): T? = if (size > 1) this[1] else null

fun render13986(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome13987 {
    data class Success(val payload: String) : Outcome13987()
    data object Failure : Outcome13987()
}

fun describe13987(outcome: Outcome13987): String = when (outcome) {
    is Outcome13987.Success -> outcome.payload
    Outcome13987.Failure -> "failed"
}
