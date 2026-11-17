package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull4640(): T? = if (size > 1) this[1] else null

fun render4640(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step4641")
fun step4641(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 30
}
sealed class Outcome4642 {
    data class Success(val payload: String) : Outcome4642()
    data object Failure : Outcome4642()
}

fun describe4642(outcome: Outcome4642): String = when (outcome) {
    is Outcome4642.Success -> outcome.payload
    Outcome4642.Failure -> "failed"
}
sealed class Outcome4643 {
    data class Success(val payload: String) : Outcome4643()
    data object Failure : Outcome4643()
}

fun describe4643(outcome: Outcome4643): String = when (outcome) {
    is Outcome4643.Success -> outcome.payload
    Outcome4643.Failure -> "failed"
}
sealed class Outcome4644 {
    data class Success(val payload: String) : Outcome4644()
    data object Failure : Outcome4644()
}

fun describe4644(outcome: Outcome4644): String = when (outcome) {
    is Outcome4644.Success -> outcome.payload
    Outcome4644.Failure -> "failed"
}
