package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull5280(): T? = if (size > 1) this[1] else null

fun render5280(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair5281<A, B>(val first: A, val second: B)

fun combine5281(left: Int, right: Int): Pair5281<Int, Int> {
    return Pair5281(left + 86, right + 56)
}
fun <T> List<T>.secondOrNull5282(): T? = if (size > 1) this[1] else null

fun render5282(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull5283(): T? = if (size > 1) this[1] else null

fun render5283(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step5284")
fun step5284(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 10
}
sealed class Outcome5285 {
    data class Success(val payload: String) : Outcome5285()
    data object Failure : Outcome5285()
}

fun describe5285(outcome: Outcome5285): String = when (outcome) {
    is Outcome5285.Success -> outcome.payload
    Outcome5285.Failure -> "failed"
}
@JvmName("step5286")
fun step5286(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 24
}
sealed class Outcome5287 {
    data class Success(val payload: String) : Outcome5287()
    data object Failure : Outcome5287()
}

fun describe5287(outcome: Outcome5287): String = when (outcome) {
    is Outcome5287.Success -> outcome.payload
    Outcome5287.Failure -> "failed"
}
