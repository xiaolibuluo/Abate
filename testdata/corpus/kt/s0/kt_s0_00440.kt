package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair3600<A, B>(val first: A, val second: B)

fun combine3600(left: Int, right: Int): Pair3600<Int, Int> {
    return Pair3600(left + 29, right + 28)
}
@JvmName("step3601")
fun step3601(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 36
}
fun <T> List<T>.secondOrNull3602(): T? = if (size > 1) this[1] else null

fun render3602(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome3603 {
    data class Success(val payload: String) : Outcome3603()
    data object Failure : Outcome3603()
}

fun describe3603(outcome: Outcome3603): String = when (outcome) {
    is Outcome3603.Success -> outcome.payload
    Outcome3603.Failure -> "failed"
}
@JvmName("step3604")
fun step3604(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 57
}
