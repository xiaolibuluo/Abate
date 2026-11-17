package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair19210<A, B>(val first: A, val second: B)

fun combine19210(left: Int, right: Int): Pair19210<Int, Int> {
    return Pair19210(left + 49, right + 36)
}
@JvmName("step19211")
fun step19211(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 56
}
@JvmName("step19212")
fun step19212(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 63
}
object Registry19213 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19213 = "registry-70"
    }
}
fun <T, R> List<T>.mapChecked19214(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair19215<A, B>(val first: A, val second: B)

fun combine19215(left: Int, right: Int): Pair19215<Int, Int> {
    return Pair19215(left + 84, right + 30)
}
