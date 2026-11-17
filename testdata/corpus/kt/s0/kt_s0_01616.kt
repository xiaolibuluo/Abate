package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step13340")
fun step13340(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 81
}
sealed class Outcome13341 {
    data class Success(val payload: String) : Outcome13341()
    data object Failure : Outcome13341()
}

fun describe13341(outcome: Outcome13341): String = when (outcome) {
    is Outcome13341.Success -> outcome.payload
    Outcome13341.Failure -> "failed"
}
fun <T> List<T>.secondOrNull13342(): T? = if (size > 1) this[1] else null

fun render13342(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome13343 {
    data class Success(val payload: String) : Outcome13343()
    data object Failure : Outcome13343()
}

fun describe13343(outcome: Outcome13343): String = when (outcome) {
    is Outcome13343.Success -> outcome.payload
    Outcome13343.Failure -> "failed"
}
