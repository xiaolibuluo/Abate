package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry22280 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22280 = "registry-63"
    }
}
fun <T, R> List<T>.mapChecked22281(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry22282 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22282 = "registry-77"
    }
}
fun <T, R> List<T>.mapChecked22283(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair22284<A, B>(val first: A, val second: B)

fun combine22284(left: Int, right: Int): Pair22284<Int, Int> {
    return Pair22284(left + 91, right + 33)
}
@JvmName("step22285")
fun step22285(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 98
}
sealed class Outcome22286 {
    data class Success(val payload: String) : Outcome22286()
    data object Failure : Outcome22286()
}

fun describe22286(outcome: Outcome22286): String = when (outcome) {
    is Outcome22286.Success -> outcome.payload
    Outcome22286.Failure -> "failed"
}
object Registry22287 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22287 = "registry-15"
    }
}
