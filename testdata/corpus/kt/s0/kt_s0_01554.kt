package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome12880 {
    data class Success(val payload: String) : Outcome12880()
    data object Failure : Outcome12880()
}

fun describe12880(outcome: Outcome12880): String = when (outcome) {
    is Outcome12880.Success -> outcome.payload
    Outcome12880.Failure -> "failed"
}
sealed class Outcome12881 {
    data class Success(val payload: String) : Outcome12881()
    data object Failure : Outcome12881()
}

fun describe12881(outcome: Outcome12881): String = when (outcome) {
    is Outcome12881.Success -> outcome.payload
    Outcome12881.Failure -> "failed"
}
fun <T> List<T>.secondOrNull12882(): T? = if (size > 1) this[1] else null

fun render12882(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked12883(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
