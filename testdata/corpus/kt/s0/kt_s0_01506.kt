package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step12460")
fun step12460(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 101
}
fun <T> List<T>.secondOrNull12461(): T? = if (size > 1) this[1] else null

fun render12461(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step12462")
fun step12462(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 18
}
fun <T, R> List<T>.mapChecked12463(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome12464 {
    data class Success(val payload: String) : Outcome12464()
    data object Failure : Outcome12464()
}

fun describe12464(outcome: Outcome12464): String = when (outcome) {
    is Outcome12464.Success -> outcome.payload
    Outcome12464.Failure -> "failed"
}
data class Pair12465<A, B>(val first: A, val second: B)

fun combine12465(left: Int, right: Int): Pair12465<Int, Int> {
    return Pair12465(left + 39, right + 63)
}
object Registry12466 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12466 = "registry-46"
    }
}
