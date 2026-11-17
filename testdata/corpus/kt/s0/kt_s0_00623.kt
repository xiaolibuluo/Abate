package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked5090(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step5091")
fun step5091(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 33
}
data class Pair5092<A, B>(val first: A, val second: B)

fun combine5092(left: Int, right: Int): Pair5092<Int, Int> {
    return Pair5092(left + 40, right + 56)
}
object Registry5093 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5093 = "registry-47"
    }
}
data class Pair5094<A, B>(val first: A, val second: B)

fun combine5094(left: Int, right: Int): Pair5094<Int, Int> {
    return Pair5094(left + 54, right + 17)
}
sealed class Outcome5095 {
    data class Success(val payload: String) : Outcome5095()
    data object Failure : Outcome5095()
}

fun describe5095(outcome: Outcome5095): String = when (outcome) {
    is Outcome5095.Success -> outcome.payload
    Outcome5095.Failure -> "failed"
}
