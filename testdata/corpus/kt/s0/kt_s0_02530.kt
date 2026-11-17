package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step20820")
fun step20820(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 8
}
fun <T, R> List<T>.mapChecked20821(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step20822")
fun step20822(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 22
}
data class Pair20823<A, B>(val first: A, val second: B)

fun combine20823(left: Int, right: Int): Pair20823<Int, Int> {
    return Pair20823(left + 29, right + 37)
}
object Registry20824 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20824 = "registry-36"
    }
}
data class Pair20825<A, B>(val first: A, val second: B)

fun combine20825(left: Int, right: Int): Pair20825<Int, Int> {
    return Pair20825(left + 43, right + 59)
}
object Registry20826 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20826 = "registry-50"
    }
}
@JvmName("step20827")
fun step20827(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 57
}
object Registry20828 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20828 = "registry-64"
    }
}
