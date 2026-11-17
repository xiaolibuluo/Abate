package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome5270 {
    data class Success(val payload: String) : Outcome5270()
    data object Failure : Outcome5270()
}

fun describe5270(outcome: Outcome5270): String = when (outcome) {
    is Outcome5270.Success -> outcome.payload
    Outcome5270.Failure -> "failed"
}
data class Pair5271<A, B>(val first: A, val second: B)

fun combine5271(left: Int, right: Int): Pair5271<Int, Int> {
    return Pair5271(left + 73, right + 49)
}
object Registry5272 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5272 = "registry-80"
    }
}
fun <T> List<T>.secondOrNull5273(): T? = if (size > 1) this[1] else null

fun render5273(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome5274 {
    data class Success(val payload: String) : Outcome5274()
    data object Failure : Outcome5274()
}

fun describe5274(outcome: Outcome5274): String = when (outcome) {
    is Outcome5274.Success -> outcome.payload
    Outcome5274.Failure -> "failed"
}
@JvmName("step5275")
fun step5275(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 101
}
fun <T, R> List<T>.mapChecked5276(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull5277(): T? = if (size > 1) this[1] else null

fun render5277(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
