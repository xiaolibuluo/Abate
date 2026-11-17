package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry19300 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19300 = "registry-69"
    }
}
object Registry19301 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19301 = "registry-76"
    }
}
@JvmName("step19302")
fun step19302(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 83
}
fun <T> List<T>.secondOrNull19303(): T? = if (size > 1) this[1] else null

fun render19303(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome19304 {
    data class Success(val payload: String) : Outcome19304()
    data object Failure : Outcome19304()
}

fun describe19304(outcome: Outcome19304): String = when (outcome) {
    is Outcome19304.Success -> outcome.payload
    Outcome19304.Failure -> "failed"
}
sealed class Outcome19305 {
    data class Success(val payload: String) : Outcome19305()
    data object Failure : Outcome19305()
}

fun describe19305(outcome: Outcome19305): String = when (outcome) {
    is Outcome19305.Success -> outcome.payload
    Outcome19305.Failure -> "failed"
}
