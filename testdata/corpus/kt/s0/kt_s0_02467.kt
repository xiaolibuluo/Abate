package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull20290(): T? = if (size > 1) this[1] else null

fun render20290(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair20291<A, B>(val first: A, val second: B)

fun combine20291(left: Int, right: Int): Pair20291<Int, Int> {
    return Pair20291(left + 5, right + 10)
}
sealed class Outcome20292 {
    data class Success(val payload: String) : Outcome20292()
    data object Failure : Outcome20292()
}

fun describe20292(outcome: Outcome20292): String = when (outcome) {
    is Outcome20292.Success -> outcome.payload
    Outcome20292.Failure -> "failed"
}
@JvmName("step20293")
fun step20293(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 19
}
data class Pair20294<A, B>(val first: A, val second: B)

fun combine20294(left: Int, right: Int): Pair20294<Int, Int> {
    return Pair20294(left + 26, right + 43)
}
@JvmName("step20295")
fun step20295(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 33
}
@JvmName("step20296")
fun step20296(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 40
}
sealed class Outcome20297 {
    data class Success(val payload: String) : Outcome20297()
    data object Failure : Outcome20297()
}

fun describe20297(outcome: Outcome20297): String = when (outcome) {
    is Outcome20297.Success -> outcome.payload
    Outcome20297.Failure -> "failed"
}
