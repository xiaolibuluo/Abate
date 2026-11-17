package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step2270")
fun step2270(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 46
}
fun <T> List<T>.secondOrNull2271(): T? = if (size > 1) this[1] else null

fun render2271(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step2272")
fun step2272(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 60
}
sealed class Outcome2273 {
    data class Success(val payload: String) : Outcome2273()
    data object Failure : Outcome2273()
}

fun describe2273(outcome: Outcome2273): String = when (outcome) {
    is Outcome2273.Success -> outcome.payload
    Outcome2273.Failure -> "failed"
}
sealed class Outcome2274 {
    data class Success(val payload: String) : Outcome2274()
    data object Failure : Outcome2274()
}

fun describe2274(outcome: Outcome2274): String = when (outcome) {
    is Outcome2274.Success -> outcome.payload
    Outcome2274.Failure -> "failed"
}
sealed class Outcome2275 {
    data class Success(val payload: String) : Outcome2275()
    data object Failure : Outcome2275()
}

fun describe2275(outcome: Outcome2275): String = when (outcome) {
    is Outcome2275.Success -> outcome.payload
    Outcome2275.Failure -> "failed"
}
