package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome14430 {
    data class Success(val payload: String) : Outcome14430()
    data object Failure : Outcome14430()
}

fun describe14430(outcome: Outcome14430): String = when (outcome) {
    is Outcome14430.Success -> outcome.payload
    Outcome14430.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked14431(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome14432 {
    data class Success(val payload: String) : Outcome14432()
    data object Failure : Outcome14432()
}

fun describe14432(outcome: Outcome14432): String = when (outcome) {
    is Outcome14432.Success -> outcome.payload
    Outcome14432.Failure -> "failed"
}
sealed class Outcome14433 {
    data class Success(val payload: String) : Outcome14433()
    data object Failure : Outcome14433()
}

fun describe14433(outcome: Outcome14433): String = when (outcome) {
    is Outcome14433.Success -> outcome.payload
    Outcome14433.Failure -> "failed"
}
@JvmName("step14434")
fun step14434(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 71
}
fun <T> List<T>.secondOrNull14435(): T? = if (size > 1) this[1] else null

fun render14435(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
