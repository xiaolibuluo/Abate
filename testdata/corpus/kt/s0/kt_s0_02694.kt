package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked22260(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome22261 {
    data class Success(val payload: String) : Outcome22261()
    data object Failure : Outcome22261()
}

fun describe22261(outcome: Outcome22261): String = when (outcome) {
    is Outcome22261.Success -> outcome.payload
    Outcome22261.Failure -> "failed"
}
@JvmName("step22262")
fun step22262(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 51
}
fun <T, R> List<T>.mapChecked22263(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step22264")
fun step22264(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 65
}
sealed class Outcome22265 {
    data class Success(val payload: String) : Outcome22265()
    data object Failure : Outcome22265()
}

fun describe22265(outcome: Outcome22265): String = when (outcome) {
    is Outcome22265.Success -> outcome.payload
    Outcome22265.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked22266(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome22267 {
    data class Success(val payload: String) : Outcome22267()
    data object Failure : Outcome22267()
}

fun describe22267(outcome: Outcome22267): String = when (outcome) {
    is Outcome22267.Success -> outcome.payload
    Outcome22267.Failure -> "failed"
}
